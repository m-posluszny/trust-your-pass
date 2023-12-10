import { Title } from '../Static/Title.component';
import { Footer } from '../Static/Footer.component';
import { OutputList } from '../Output/Output.component';
import { PassInput } from '../Input/Input.component';
import { useGetPassword } from '../Input/Input.hook';
import { ErrorComponent } from '../Error/Error.component';

export const MainView = () => {
  const { sendPassword, isLoading, error, createError } = useGetPassword();
  return (
    <>
      <main class="flex my-auto flex-col  overflow-y-visible overflow-x-hidden h-[100vh]">
        <Title />
        <PassInput sendInput={sendPassword} isLoading={isLoading} />

        {error ? <ErrorComponent error={error} /> : createError ? <ErrorComponent error={createError} /> :
          <OutputList />
        }

      </main>
      <Footer />
    </>
  )
}