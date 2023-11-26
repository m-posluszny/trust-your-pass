const Background = "bg-[radial-gradient(ellipse_at_top_right,_var(--tw-gradient-stops))] from-slate-900 via-purple-800 to-slate-900";

export const Container = ({ children }) => (
    <div className={`flex flex-col ${Background}`} style={{ height: "100vh" }}>
        {children}
    </div>
)