import { OutputColor, OutputIcon, OutputSuccess, OutputLoading, OutputError } from "./Output.type";
import {Circle} from "../Misc/Circle.component";
import {ExpandButton} from "../Misc/Buttons.component";
import { useState } from "react";

export const OutputTile = ({ type, message }) => <>
          <div class={ `flex w-full my-5 p-3 rounded-xl ${OutputColor(type)}  ${type === OutputLoading && "animate-pulse"}` }>
            <OutputIcon type={type} />
            <p className='text-xl text-center my-auto ms-5'>
                {message}
            </p>
          </div>
</>

export const PassDetails = ({ output }) => {
    return <div class="bg-opacity-70 bg-white rounded-2xl p-2 text-xl flex w-fit">
        <Circle> 1 </Circle>
        <div className="py-1 mx-3 text-center text-xl">
            12:34 | 26.11.2023 
        </div>
    </div>

 }


export const OutputSubList = ({ output, expanded }) => { 
    const [ isExpanded, setExpanded ] = useState(expanded)
    return <div class="bg-opacity-20 bg-white rounded-2xl p-3 text-xl mb-3" >
        <div className="flex">
        <PassDetails/>
        <ExpandButton className="ms-auto p-2 px-4 bg-white bg-opacity-70 rounded-xl" isExpanded={isExpanded} setExpanded={setExpanded}/>
        </div>
        {isExpanded && <>
        <OutputTile type={ OutputSuccess } message="Password length validation successful"/>
        <OutputTile type={OutputError} message="Too many repeated characters"/>
        <OutputTile type={ OutputLoading}  message="Model check in progress..."/>
        </>}
    </div>


}

export const OutputList = ({ useOutputs }) => {

    return <div class=" p-5 rounded-2xl  overflow-y-visible overflow-x-hidden  mx-auto w-2/5" style={{ "maxHeight": "50vh", "minWidth": "400pt" }}>
        <OutputSubList expanded/>
        <OutputSubList/>
    </div>

 }
