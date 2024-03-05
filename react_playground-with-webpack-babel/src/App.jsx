import React from "react";
import "./style/app.css";
import Greeting from "./components/Greeting";

const App = () => {
  return (
    <div>
      <h1 className="sample">Hello React with Webpack!!!!</h1>
      <Greeting />
    </div>
  );
};

export default App;
