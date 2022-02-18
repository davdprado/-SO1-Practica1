import "./App.css";
import Calc from "./Components/Calc";
import Tabla from "./Components/Tabla";

function App() {
  return (
    <div className="App">
      <h1>Calculadora</h1>
      <Calc />
      <br></br>
      <Tabla />
    </div>
  );
}

export default App;
