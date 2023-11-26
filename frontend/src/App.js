import './App.css';
import { FaExclamationCircle, FaQuestionCircle, FaCheckCircle } from "react-icons/fa";

const Background = "bg-[radial-gradient(ellipse_at_top_right,_var(--tw-gradient-stops))] from-slate-900 via-purple-900 to-slate-900";

function App() {
  return (
    <div className={`flex flex-col ${Background}` } style={{ height: "100vh" }}>
      <main class="flex my-auto flex-col">
        <h1 class="text-white text-6xl text-center mt-0">
          Trust Pass
        </h1>
        <h3 class="text-white text text-center mb-36 mt-0">
          Check if your password is unique enough
        </h3>
        <form class="flex flex-col">
          <input type="password" id="last_name" class="px-5 py-2 w-25 mx-auto bg-purple-800 hover:bg-purple-700 text-4xl rounded-3xl text-white placeholder-slate-200 text-center shadow-md" placeholder="Check your password" required>

          </input>
          <button class="text-white bg-green-500  hover:bg-green-400 w-min mx-auto px-5 py-2 m-5 text-lg rounded-xl shadow-md"> 
            Submit
          </button>
        </form>
        <div class="overflow-y-visible overflow-x-hidden  mx-auto w-2/5" style={{"maxHeight":"50vh"}}>
          <div class="flex w-11/12 m-5 p-3 rounded-xl  bg-green-400">
            <FaCheckCircle class="inline-block text-5xl text-gray-600"/>
            <p className='text-xl text-center my-auto ms-5'>
            Password length validation successful
            </p>
          </div>
          <div class="flex w-11/12 m-5 p-3 rounded-xl  bg-red-400">
            <FaExclamationCircle class="inline-block text-gray-600 text-5xl"/>
            <p className='text-xl text-center my-auto ms-5'>
            Too many repeated characters
            </p>
          </div>
          <div class="flex w-11/12 m-5 p-3 rounded-xl  bg-yellow-400 animate-pulse">
            <FaQuestionCircle class="inline-block text-gray-600 text-5xl"/>
            <p className='text-xl text-center my-auto ms-5'>
            Model check in progress...
            </p>
          </div>

        </div>
      </main>
      <footer className="mb-2 mx-3 mt-auto text-center text-white flex flex-row">
        <p className="me-auto">
          Copyright 2023
        </p>
        <a className="ms-auto" href="https://github.com/m-posluszny/trust-your-pass/tree/feature/gateway-service">
          Github
        </a>
        <a className="ms-5" href="/about">
          About
        </a>
        <a className="ms-5" href="/api/docs">
          API Docs
        </a>

      </footer>
    </div>
  );
}

export default App;
