import { createAuth0 } from "@auth0/auth0-vue";
import { createApp } from "vue";
import Vue3Toastify, { type ToastContainerOptions } from "vue3-toastify";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import { aliases, mdi } from "vuetify/iconsets/mdi";

import App from "./App.vue";
import { vCurrency } from "./directives/vCurrency";
import routes from "./routes";
import "vue3-toastify/dist/index.css";
import "vuetify/styles";
import "./style.css";
import "@mdi/font/css/materialdesignicons.css"; // Ensure you are using css-loader

const vuetify = createVuetify({
  components,
  directives,
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: {
      mdi
    }
  }
});

async function loadConfig() {
  try {
    const response = await fetch("/config.json");
    const config = await response.json();
    return config;
  } catch (error) {
    console.error("Failed to load configuration:", error);
    throw new Error("Configuration loading failed");
  }
}

loadConfig()
  .then((config) => {
    createApp(App)
      .use(routes)
      .use(vuetify)
      .use(Vue3Toastify, {
        clearOnUrlChange: false
      } as ToastContainerOptions)
      .use(
        createAuth0(
          {
            domain: config.auth0.domain,
            clientId: config.auth0.clientId,
            authorizeTimeoutInSeconds: 2,
            authorizationParams: {
              max_age: "1d",
              audience: config.auth0.audience,
              redirect_uri: window.location.origin
            },
            cacheLocation: "localstorage"
          },
          {
            // skipRedirectCallback: true
          }
        )
      )
      .directive("currency", vCurrency)
      .mount("#app");
  })
  .catch((error) => {
    console.error("App initialization failed:", error);
  });
