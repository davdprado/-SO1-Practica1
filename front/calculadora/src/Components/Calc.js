import React, { useState, useEffect } from "react";

export default function Calc() {
  const [conta, setConta] = useState("");
  let Num1 = 0;
  let Num2 = 0;
  let Simbolo = "";
  let Resultado = 0;

  const operar = () => {
    const operacion = conta.split(" ");
    Num1 = parseFloat(operacion[0]);
    Num2 = parseFloat(operacion[2]);
    Simbolo = operacion[1];
    switch (Simbolo) {
      case "+":
        Resultado = Num1 + Num2;
        break;
      case "-":
        Resultado = Num1 - Num2;
        break;
      case "*":
        Resultado = Num1 * Num2;
        break;
      case "/":
        Resultado = Num1 / Num2;
        break;
    }
    setConta(Resultado.toString());
  };
  return (
    <>
      <input type={Text} value={conta} />
      <br></br>
      <button onClick={() => setConta(conta + "7")}>7</button>
      <button onClick={() => setConta(conta + "8")}>8</button>
      <button onClick={() => setConta(conta + "9")}>9</button>
      <button onClick={() => setConta("")}>AC</button>
      <br></br>
      <button onClick={() => setConta(conta + "4")}>4</button>
      <button onClick={() => setConta(conta + "5")}>5</button>
      <button onClick={() => setConta(conta + "6")}>6</button>
      <button onClick={() => setConta(conta + " * ")}>X</button>
      <button onClick={() => setConta(conta + " / ")}>/</button>
      <br></br>
      <button onClick={() => setConta(conta + "1")}>1</button>
      <button onClick={() => setConta(conta + "2")}>2</button>
      <button onClick={() => setConta(conta + "3")}>3</button>
      <button onClick={() => setConta(conta + " + ")}>+</button>
      <button onClick={() => setConta(conta + " - ")}>-</button>
      <br></br>
      <button onClick={() => setConta(conta + "0")}>0</button>
      <button onClick={() => setConta(conta + ".")}>.</button>
      <button onClick={() => setConta(operar)}>=</button>
      <br></br>
      <br></br>
    </>
  );
}
