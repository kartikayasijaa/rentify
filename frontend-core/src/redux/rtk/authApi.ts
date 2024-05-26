import { createApi } from "@reduxjs/toolkit/query/react";
import { baseApiQuery } from ".";
import { LoginRequestType } from "../../modules/Login/types";

export const authApi = createApi({
    reducerPath: "authApi",
    baseQuery: baseApiQuery,
    endpoints: (builder) => ({
        register: builder.mutation({
            query: (user) => ({
                url: "auth/register",
                method: "POST",
                body: user
            })
        }),
        login: builder.mutation({
            query: (credentials: LoginRequestType) => ({
                url: "auth/login",
                method: "POST",
                body: credentials
            })
        }),
        refresh: builder.query({
            query: () => ({
                url: "auth/refresh",
                method: "GET"
            })
        })
    })
});

export const { useRegisterMutation, useLoginMutation, useRefreshQuery } =
    authApi;
