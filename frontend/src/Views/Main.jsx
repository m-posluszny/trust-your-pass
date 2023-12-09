import { Title } from '../Static/Title.component';
import { Footer } from '../Static/Footer.component';
import { OutputList } from '../Output/Output.component';
import { PassInput } from '../Input/Input.component';

export const MainView = () => ( 
    <>
      <main class="flex my-auto flex-col  overflow-y-visible overflow-x-hidden h-[100vh]">
        <Title/>
        <PassInput/>
        <OutputList/>
      </main>
      <Footer/>
    </>
 )