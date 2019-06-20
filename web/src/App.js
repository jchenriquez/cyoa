import React from 'react';
import axios from 'axios';
import BookList from './components/booklist';
import { BrowserRouter as Router, Route} from "react-router-dom";
import {createBrowserHistory} from 'history';
import './App.css';
import StoryArc from "./components/storyarc";

axios.defaults.baseURL = `http://${process.env.SERVER_HOST}:${process.env.SERVER_PORT}`;
axios.defaults.headers.post['Content-Type'] = 'application/json';

function App() {
  return (
      <Router history={createBrowserHistory()}>
          <div>
              <Route exact path={'/'} component={BookList}/>
              <Route path={'/:book/:story'} component={StoryArc}/>
          </div>
      {/*<header className="App-header">*/}
      {/*  <img src={logo} className="App-logo" alt="logo" />*/}
      {/*</header>*/}

        </Router>

  );
}

export default App;
