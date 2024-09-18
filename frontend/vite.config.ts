import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import fs from "fs";
import path from "path";
import { config } from "dotenv";

config();
const dynamicConfigPlugin = () => ({
  name: "dynamic-config",
  configureServer(server) {
    server.middlewares.use((req, res, next) => {
      if (req.url.endsWith("config.json")) {
        const configPath = path.resolve("public/config.json");
        fs.readFile(configPath, "utf8", (err, data) => {
          if (err) {
            console.error("Error reading config.json:", err);
            res.statusCode = 500;
            res.end("Internal Server Error");
            return;
          }

          let modifiedConfig = data
            .replace(/\$AUTH0_APP_DOMAIN/g, process.env.AUTH0_APP_DOMAIN || "")
            .replace(
              /\$AUTH0_APP_CLIENT_ID/g,
              process.env.AUTH0_APP_CLIENT_ID || ""
            )
            .replace(
              /\$AUTH0_APP_AUDIENCE/g,
              process.env.AUTH0_APP_AUDIENCE || ""
            );

          res.setHeader("Content-Type", "application/json");
          res.end(modifiedConfig);
        });
      } else {
        next();
      }
    });
  }
});

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), dynamicConfigPlugin()],
  resolve: {
    alias: [
      {
        find: "@",
        replacement: fileURLToPath(new URL("./src", import.meta.url))
      }
    ]
  },
  define: {
    APP_VERSION: JSON.stringify(process.env.npm_package_version)
  }
});
