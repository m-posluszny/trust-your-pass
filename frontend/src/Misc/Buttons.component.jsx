export const ExpandButton = ({className, isExpanded, setExpanded }) => ( 
    <button className={className} onClick={()=>setExpanded(!isExpanded)}>
      {isExpanded ? "▲" : "▼"}

    </button>
  ) 