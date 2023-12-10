import { OutputColor, OutputIcon, OutputSuccess, OutputLoading, OutputError } from "./Output.type";
import { Circle } from "../Misc/Circle.component";
import { ExpandButton } from "../Misc/Buttons.component";
import { useState } from "react";
import {
    useResults, useResult
} from "./Output.hook";

export const OutputTile = ({ type, message }) => <>
    <div class={`flex w-full my-5 p-3 rounded-xl ${OutputColor(type)}  ${type === OutputLoading && "animate-pulse"}`}>
        <OutputIcon type={type} />
        <p className='text-xl  my-auto ms-5'>
            {message}
        </p>
    </div>
</>

export const PassHeader = ({ index, date }) => {
    let d = new Date(date)
    return <div class="bg-opacity-70 bg-white rounded-2xl p-2 text-md flex w-fit">
        <Circle> {index} </Circle>
        <div className="py-1 mx-3 text-center">
            {d.toLocaleString().replace(",", " |")}
        </div>
    </div>

}


export const OutputSubList = ({ index, id, expanded }) => {
    const [isExpanded, setExpanded] = useState(expanded)
    const { output, error } = useResult(id)
    console.log(output)
    if (error || !output) {
        return <></>
    }

    return <div class="bg-opacity-20 bg-white rounded-2xl p-3 text-xl mb-3" >
        <div className="flex">
            <PassHeader date={output.date} index={index} />
            <ExpandButton className="ms-auto   bg-white bg-opacity-70 rounded-xl" isExpanded={isExpanded} setExpanded={setExpanded} />
        </div>
        <div className={`${isExpanded ? 'h-full' : 'h-0'
            } transition-[height]  duration-100 overflow-hidden w-full`}>
            {output.strength > -1 ?
                <OutputTile type={OutputSuccess} message={`Password strength: ${output.strength}`} /> :
                !output.isProcessed ?
                    < OutputTile type={OutputLoading} message="Password strength check in progress..." /> : <></>
            }
            {output?.preconditions?.map((precondition) =>
                <OutputTile type={precondition.isSatisfied ? OutputSuccess : OutputError} message={precondition.condition} />
            )}

        </div>
    </div>


}

export const OutputList = () => {
    const { results } = useResults()
    return <div class=" p-5 rounded-2xl mx-auto w-2/5" style={{ "minWidth": "400pt" }}>
        {results.map(((result, i) =>
            <OutputSubList id={result.id} index={results.length - i} output={result} expanded={i === 0} />
        ))}
    </div>

}
