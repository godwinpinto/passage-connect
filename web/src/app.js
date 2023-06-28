import React from "react";
import { Routes, Route } from 'react-router-dom';

import Home from "./views/Home";
import Dashboard from "./views/Dashboard";
import Banner from "./components/banner";
import styles from './styles/App.module.css';

function App() {
  return (
      <div>
            <Banner/>
            <div className={styles.mainContainer}>
                <Routes>
                    <Route path="/" element={<Home/>}></Route>
                    <Route path="/dashboard" element={<Dashboard/>}></Route>
                </Routes>
            </div>
            <div className={styles.footer}>
                This is a demo application for Hackathon, please do not login.<br/><br/>      
            </div>
            <div className={styles.footer}>
                Learn more with our <a href="https://github.com/godwinpinto/passage-connect" target="_blank" rel="noopener">Github</a>.      
            </div>
      </div>
  );
}

export default App;
