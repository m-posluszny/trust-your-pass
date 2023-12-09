import { useEffect, useState } from "react";
import { useCore} from "../Api/core.hook"
import { useResults } from "../Output/Output.hook";
import { post } from "../Api/core";

export const PASSWORD_STATUS = (uid) => `api/v1/passwords/${uid}`;

export const useGetPassword = () => { 
    const [uid, setUid] = useState(null);
    const [isProcessing, setIsProcessing] = useState(false);
    const [loading, setLoading] = useState(false);
    const {editResult, resultsMap} = useResults();

    const { data, error } = useCore(
        PASSWORD_STATUS(uid),
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
        error,
        loading,
        sendPassword: (password) => { 
            setLoading(true);
            post('http://localhost:5000/api/password', { password }).then((res) => {
                setUid(res.uid);
                setIsProcessing(res.isProcessing);
                editResult(res);
                setLoading(false);
            })
        } 
    }
} 