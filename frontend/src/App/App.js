import { BrowserRouter, Route, Routes } from "react-router-dom";
import { MainView} from "../Views/Main";
import {AboutView } from "../Views/About"
import { Container } from "../Static/Container.component";


function App() {
  return (
    <Container>
      <BrowserRouter>
        <Routes>
          <Route exact path="/about" element={<AboutView/> }/>
          <Route path="/" element={ <MainView/>}/>
      </Routes>
      </BrowserRouter>
    </Container>
  );
}

export default App;
