import { Directive } from "vue";

const vCurrency: Directive<HTMLInputElement> = {
  beforeMount(el, _, vnode) {
    const inputElement =
      el.tagName === "INPUT" ? el : el.querySelector("input");

    if (!inputElement) return;

    inputElement.classList.add("v-currency");

    const formatCurrency = (number: string) => {
      const parsed = parseFloat(number.replace(/[^\d.-]/g, ""));
      return new Intl.NumberFormat("en-AU", {
        style: "currency",
        currency: "AUD"
      })
        .format(parsed)
        .replace("$", "");
    };

    const updateValue = () => {
      const formattedValue = formatCurrency(inputElement.value);
      if (inputElement.value !== formattedValue) {
        inputElement.value = formattedValue;
      }
    };

    inputElement.addEventListener("focus", () => {
      // On focus, convert the formatted value back to a plain number
      const numberValue = parseFloat(
        inputElement.value.replace(/[^\d.-]/g, "")
      );

      if (!isNaN(numberValue)) {
        inputElement.value = numberValue.toString();
      }
    });

    inputElement.addEventListener("blur", updateValue);

    // Handle model updates from Vue (e.g., v-model changes)
    vnode.props ||= {};
    const originalInputHandler =
      vnode.props["onUpdate:modelValue"] || (() => {});
    vnode.props["onUpdate:modelValue"] = (...args: any[]) => {
      originalInputHandler(...args);
      updateValue(); // Ensure the input is formatted
    };

    // Initial formatting
    updateValue();
  },
  updated(el) {
    // Ensure the formatting is reapplied if the element updates
    const inputElement =
      el.tagName === "INPUT" ? el : el.querySelector("input");

    if (!inputElement || document.activeElement === inputElement) return;

    inputElement.dispatchEvent(new Event("blur"));
  }
};

export { vCurrency };
