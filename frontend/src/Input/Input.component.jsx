
export const PassInput = () => {
        return <form class="flex flex-col">
          <input type="password" id="last_name" class="px-5 py-2 w-25 mx-auto bg-purple-800 hover:bg-purple-700 text-4xl rounded-3xl text-white placeholder-slate-200 text-center shadow-md" placeholder="Check your password" required>

          </input>
          <button class="text-white bg-green-500  hover:bg-green-400 w-min mx-auto px-5 py-2 m-5 text-lg rounded-xl shadow-md"> 
            Submit
          </button>
        </form>
}