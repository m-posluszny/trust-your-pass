import { useEffect, useState } from "react";
import { useCore} from "../Api/core.hook"
import { useLocalStorage } from "@uidotdev/usehooks";

export const PASSWORDS = (uid) => `api/v1/passwords/${uid}`;

export const useResults = () => {
    const [results, setResults] = useLocalStorage("results", {});

    const editResult = (result) => { 
        results[result.id] = result;
        if (!results[result.id].date) {
            results[result.id].date = (new Date()).toString();} ;
        setResults(results);
    }
    
    return {
        resultsMap: results,
        results: Object.values(results).sort((a, b) => Date.parse(b.date) - Date.parse(a.date)),
        editResult,
    }
}

export const useResult = (id) => {
    const {editResult, resultsMap} = useResults();
    const [output, setOutput] = useState(resultsMap[id]);

    const { data, error } = useCore(
        PASSWORDS(id),
        {},
        typeof id === "string" && id !== "",
        {
          refreshInterval: output?.isProcessing ? 10 : 0,
          revalidateIfStale: false,
          revalidateOnFocus: false,
          revalidateOnReconnect: false,
        }
      );

    console.log(data)
    useEffect(() => {
        setOutput(resultsMap[id]);
    },[])
    
    useEffect(() => {
        if (data && resultsMap[id] !== data) 
        {
            editResult(data);
        }
    },[])

    return {
        output,
        error,
    }


}