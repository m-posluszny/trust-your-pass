import { useState } from "react";
import { useResults } from "../Output/Output.hook";
import { post } from "../Api/core";

export const PASSWORD = () => `api/v1/passwords`;

export const useGetPassword = () => { 
    const [loading, setLoading] = useState(false);
    const {editResult} = useResults();
    const [createError, setError] = useState(null);


    return {
        createError,
        loading,
        sendPassword: (password) => { 
            setLoading(true);
            post(PASSWORD(), { password }).then((res) => {
                editResult(res);
                setLoading(false);
            }).catch((err) => {
                setError(err);
                setLoading(false);
            })
        } 
    }
} 