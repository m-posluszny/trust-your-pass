
export const PassInput = ({ sendInput }) => {
  const inputFromForm = (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    return formData.get('password')
  }
  return <form class="flex flex-col mt-10" onSubmit={e => sendInput(inputFromForm(e))}>
    <input type="password" name="password" class="px-5 py-2 w-25 mx-auto bg-white bg-opacity-20 hover:bg-purple-700 text-4xl rounded-3xl text-white placeholder-slate-200 text-center shadow-md" placeholder="Check your password" required>
    </input>
    <button class="text-white bg-green-500  hover:bg-green-400 w-min mx-auto px-5 py-2 m-5 text-lg rounded-xl shadow-md">
      Submit
    </button >
  </form >
}