import React from 'react';
import { Switch, Route } from 'react-router-dom';
import { Provider } from 'react-redux';
import { CookiesProvider } from 'react-cookie';
import store from './redux/store';
import './App.css';
import Authentication from './connection/Authentication';
import Register from './connection/Register';
import Home from './home/Home';
import NotFound from './common/NotFound';
import Logout from './common/Logout';
import SourceView from './common/SourceView';

const App = () => {
  document.body.style = 'background: darkgray';
  document.title = 'Squawk Project';
  return (
    <div className="App">
      <Provider store={store}>
        <CookiesProvider>
          <Switch>
            <Route exact path="/" render={() => <Home />} />
            <Route exact path="/auth" render={() => <Authentication />} />
            <Route exact path="/register" render={() => <Register />} />
            <Route exact path="/logout" render={() => <Logout />} />
            <Route exact path="/404" render={() => <NotFound />} />
            <Route exact path="/:sourceid" render={() => <SourceView />} />
          </Switch>
        </CookiesProvider>
      </Provider>
    </div>
  );
};


export default App;
