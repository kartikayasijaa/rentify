import { createApi } from "@reduxjs/toolkit/query/react";
import { baseApiQuery } from ".";

export const propertyApi = createApi({
    reducerPath: "propertyApi",
    baseQuery: baseApiQuery,
    endpoints: (builder) => ({
        getProperties: builder.query({
            query: () => "property"
        }),
        createProperty: builder.mutation({
            query: (property) => ({
                url: "property",
                method: "POST",
                body: property
            })
        })
    })
});

export const { useGetPropertiesQuery, useCreatePropertyMutation } = propertyApi;
