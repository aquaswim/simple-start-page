import { ref } from "vue";
import type { IErrorState, IResponse } from "./data";

export const fetchApi = async <T>(method: string, url: string, body: any | null, token?: string) : Promise<IResponse<T> | null> => {
    const headers: { [key: string]: string; } = {
        'Accept': 'application/json',
    }
    const reqInit : RequestInit = {method: method}
    if (body) {
        reqInit.body = JSON.stringify(body)
        headers['Content-Type'] = 'application/json'
    }
    if (token) {
        headers['Authorization'] = token
    }
    reqInit.headers = headers
    const resp = await fetch(url, reqInit);
    return await resp.json()
}

export const useFetch = <T = any>(method: string, url: string) => {
    const isLoading = ref(false)
    const data = ref<T | null>(null)
    const error = ref<IErrorState | null> (null)
    const execute = async(body?: any, token?: string) => {
        try {
            isLoading.value = true
            data.value = null
            error.value = null
            const resp = await fetchApi<T>(method, url, body, token)
            
            if (resp && resp.code == 0) {
                data.value = resp.data as any
            } else {
                // todo get response message
                throw new Error("Unexpected response body")
            }
        } catch(err) {
            error.value = {
                message: `Error "${err}" when fetching ${method} ${url}`,
            }
            data.value = null
        } finally {
            isLoading.value = false;
        }
    }

    return {isLoading, data, error, execute}
}