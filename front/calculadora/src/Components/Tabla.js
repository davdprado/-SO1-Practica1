import React, { useState, useEffect } from "react";

function useDatos() {
  const [operaciones, setOperaciones] = useState([]);
  useEffect(() => {
    fetch("http://localhost:8080/obtener")
      .then((response) => response.json())
      .then((datos) => {
        setOperaciones(datos);
      });
  }, []);
  return operaciones;
}

export default function Tabla() {
  const operaciones = useDatos();
  return (
    <>
      <h1>Registros</h1>
      <table>
        <thead>
          <tr>
            <th scope="col">Numero1</th>
            <th scope="col">Numero1</th>
            <th scope="col">Operacion</th>
            <th scope="col">Resultado</th>
            <th scope="col">Fecha</th>
          </tr>
        </thead>
        <tbody>
          {operaciones.map((item) => (
            <tr key={item._id}>
              <td>{item.numero1}</td>
              <td>{item.numero2}</td>
              <td>{item.operacion}</td>
              <td>{item.resultado}</td>
              <td>{item.fecha}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
}
