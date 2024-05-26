import { fetchBaseQuery } from "@reduxjs/toolkit/query/react";

const getToken = (): string => {
    // Implement logic to get the token, e.g., from local storage
    return "";
};

const prepareHeaders = (headers: Headers): Headers => {
    const token = getToken();
    if (token) {
        headers.set("Authorization", `Bearer ${token}`);
    }
    return headers;
};

const generateBaseURL = (prefix: string) => {
    return `http://localhost:8080/${prefix}`;
};

export const buildApiQuery = (prefix: string) => {
    return (args: any, api: any, extraOptions: any) => {
        const baseURL = generateBaseURL(prefix);
        return fetchBaseQuery({
            baseUrl: baseURL,
            prepareHeaders: prepareHeaders,
            credentials: "include",
            mode: "cors",
        })(args, api, extraOptions);
    };
};

export const baseApiQuery = buildApiQuery("api");