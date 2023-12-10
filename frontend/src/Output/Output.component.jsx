import { OutputColor, OutputIcon, OutputSuccess, OutputLoading, OutputError } from "./Output.type";
import { Circle } from "../Misc/Circle.component";
import { ExpandButton } from "../Misc/Buttons.component";
import { useState } from "react";
import { useResults } from "./Output.hook";

export const OutputTile = ({ type, message }) => <>
    <div class={`flex w-full my-5 p-3 rounded-xl ${OutputColor(type)}  ${type === OutputLoading && "animate-pulse"}`}>
        <OutputIcon type={type} />
        <p className='text-xl text-center my-auto ms-5'>
            {message}
        </p>
    </div>
</>

export const PassHeader = ({ output }) => {
    return <div class="bg-opacity-70 bg-white rounded-2xl p-2 text-md flex w-fit">
        <Circle> 1 </Circle>
        <div className="py-1 mx-3 text-center">
            {output.date.toLocaleString().replace(",", " |")}
        </div>
    </div>

}


export const OutputSubList = ({ output, expanded }) => {
    const [isExpanded, setExpanded] = useState(expanded)
    return <div class="bg-opacity-20 bg-white rounded-2xl p-3 text-xl mb-3" >
        <div className="flex">
            <PassHeader output={output} />
            <ExpandButton className="ms-auto   bg-white bg-opacity-70 rounded-xl" isExpanded={isExpanded} setExpanded={setExpanded} />
        </div>
        <div className={`${isExpanded ? 'h-72' : 'h-0'
            } transition-[height]  duration-100 overflow-hidden w-full`}>
            {output?.preconditions?.map((precondition) =>
                <OutputTile type={precondition.isValid ? OutputSuccess : OutputError} message={precondition.message} />
            )}

            <OutputTile type={OutputSuccess} message="Password length validation successful" />
            <OutputTile type={OutputError} message="Too many repeated characters" />
            {output.strength > -1 ?
                <OutputTile type={OutputSuccess} message={`AI Model Password strength: ${output.strength}`} /> :
                <OutputTile type={OutputLoading} message="Model check in progress..." />
            }
        </div>
    </div>


}

export const OutputList = () => {
    const { results } = useResults()
    return <div class=" p-5 rounded-2xl mx-auto w-2/5" style={{ "minWidth": "400pt" }}>
        {results.map(((result, i) =>
            <OutputSubList output={result} expanded={i === 0} />
        ))}
    </div>

}
