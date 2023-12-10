import { useEffect, useState } from "react";
import { useCore} from "../Api/core.hook"
import { useResults } from "../Output/Output.hook";
import { post } from "../Api/core";

export const PASSWORDS = (uid) => `api/v/passwords/${uid}`;
export const PASSWORD = () => `api/v/passwords`;

export const useGetPassword = () => { 
    const [uid, setUid] = useState(null);
    const [isProcessing, setIsProcessing] = useState(false);
    const [loading, setLoading] = useState(false);
    const {editResult, resultsMap} = useResults();
    const [createError, setError] = useState(null);

    const { data, error } = useCore(
        PASSWORDS(uid),
        {},
        typeof uid === "string" && uid !== "",
        {
          refreshInterval: isProcessing ? 10 : 0,
          revalidateIfStale: false,
          revalidateOnFocus: false,
          revalidateOnReconnect: false,
        }
      );

    useEffect(() => {
        if (data || resultsMap[uid] !== data) 
            editResult(data);
        
    })

    return {
        createError,
        error,
        loading,
        sendPassword: (password) => { 
            setLoading(true);
            post(PASSWORD(), { password }).then((res) => {
                setUid(res.uid);
                setIsProcessing(res.isProcessing);
                editResult(res);
                setLoading(false);
            }).catch((err) => {
                setError(err);
                setLoading(false);
                setIsProcessing(false);
            })
        } 
    }
} 