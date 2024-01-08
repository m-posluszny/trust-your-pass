import { useEffect, useState } from "react";
import { useCore} from "../Api/core.hook"
import { useLocalStorage } from "@uidotdev/usehooks";
import { OutputError, OutputLoading, OutputSuccess } from "./Output.type";

export const PASSWORDS = (uid) => `api/v1/passwords/${uid}`;

export const useResults = () => {
    const [results, setResults] = useLocalStorage("results", []);

    const addResult = (result, metadata) => {
        result.date =  new Date().toString();
        result.meta = metadata
        if (results.length > 10) {
            setResults([...results.slice(1), result])
        } else (
        setResults([...results, result])

        )

    }

    const editResult = (i, result) => { 
        result.meta = results[i].meta;
        results[i] = result;
        setResults(results);
    }
    
    return {
        results:results.sort((a, b) => Date.parse(b.date) - Date.parse(a.date)),
        editResult,
        addResult
    }
}

export const useResult = (i) => {
    const {editResult,  results} = useResults();
    const [output, setOutput] = useState(results[i]);
    let state = OutputError
    if (output?.IsProcessing) {
        state = OutputLoading
    } else if (output?.preconditions?.every && output?.preconditions?.every((prec => prec.isSatisfied))) {
        state = OutputSuccess
    }

    const { data, error } = useCore(
        PASSWORDS(i),
        {},
        typeof i === "string" && i !== "",
        {
          refreshInterval: output?.isProcessing ? 10 : 0,
          revalidateIfStale: false,
          revalidateOnFocus: false,
          revalidateOnReconnect: false,
        }
      );

    useEffect(() => {
        setOutput(results[i]);
    },[i, results])
    
    useEffect(() => {
        if (data && results[i] !== data) 
        {
            editResult(i, data);
        }
    },[])

    return {
        state,
        output,
        error,
    }


}