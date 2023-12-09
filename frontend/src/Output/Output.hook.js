import { useLocalStorage } from "@uidotdev/usehooks";

export const useResults = () => {
    const [results, setResults] = useLocalStorage("results", {});
    const editResult = (result) => { 
        results[result.id] = result;
        if (!results[result.id].date) {
            results[result.id].date = new Date();
        } ;
        setResults(results);
    }
    return {
        resultsMap: results,
        results: Object.values(results).sort((a, b) => b.date - a.date),
        editResult,
    }
}