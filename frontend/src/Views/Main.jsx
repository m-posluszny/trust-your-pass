import { Title } from '../Static/Title.component';
import { Footer } from '../Static/Footer.component';
import { OutputList } from '../Output/Output.component';
import { PassInput } from '../Input/Input.component';
import { useGetPassword } from '../Input/Input.hook';

export const MainView = () => {
  const { sendPassword, isLoading } = useGetPassword();
  return (
    <>
      <main class="flex my-auto flex-col  overflow-y-visible overflow-x-hidden h-[100vh]">
        <Title />
        <PassInput sendInput={sendPassword} isLoading={isLoading} />
        <OutputList />
      </main>
      <Footer />
    </>
  )
}