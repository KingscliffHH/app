import { AxiosError } from "axios";
import { UnwrapRef, ref } from "vue";

type QueryFunction<TResult> = () => Promise<TResult>;

type QueryProps<TResult> = {
  queryFn: QueryFunction<TResult>;
  onError?: (error: any) => void;
};
export const useQuery = <TResult>(props: QueryProps<TResult>) => {
  const isLoading = ref(true);
  const data = ref<TResult>(null as TResult);
  const cachedData = ref<TResult | null>(null);
  const error = ref<any>(null);

  const fetch = async () => {
    isLoading.value = true;
    try {
      const response = await props.queryFn();
      data.value = response as UnwrapRef<TResult>;
      cachedData.value = JSON.parse(JSON.stringify(data.value));
      // props.onSuccess && props.onSuccess(res);
    } catch (e) {
      if (e instanceof AxiosError) {
        error.value = e.response?.data;
      } else {
        error.value = e;
      }

      props.onError && props.onError(error.value);
    } finally {
      isLoading.value = false;
    }

    // fetch();
    // isLoading.value = true;
    // props.queryFn().then((response: TResult) => {
    //   data.value = response as UnwrapRef<TResult>;
    //   cachedData.value = JSON.parse(JSON.stringify(data.value));
    //   isLoading.value = false;
    // });
  };

  fetch();
  const reset = () => {
    data.value = JSON.parse(
      JSON.stringify(cachedData.value)
    ) as UnwrapRef<TResult>;
  };

  return { isLoading, refetch: fetch, reset, data };
};

// Define the function type with explicit argument and return types
type MutationFunction<TArg extends any[], TResult> = (
  ...args: TArg
) => Promise<TResult>;

type MutationProps<TArg extends any[], TResult> = {
  mutationFn: MutationFunction<TArg, TResult>;
  onSuccess?: (data: TResult) => void;
  onError?: (error: any) => void;
};

export const useMutation = <TArg extends any[], TResult>(
  props: MutationProps<TArg, TResult>
) => {
  const isLoading = ref(false);
  // const data = ref<TResult | null>(null);
  const error = ref<any>(null);

  const mutate = async (...args: TArg) => {
    isLoading.value = true;
    try {
      const res = await props.mutationFn(...args);
      props.onSuccess && props.onSuccess(res);
    } catch (e) {
      if (e instanceof AxiosError) {
        if (e.response?.data instanceof Object) {
          error.value = e.response?.data;
        } else {
          error.value = {
            error: e.response?.data
          };
        }
      } else {
        error.value = e;
      }

      props.onError && props.onError(error.value);
    } finally {
      isLoading.value = false;
    }
  };

  return { isLoading, error, mutate };
};
