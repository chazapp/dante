import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import { withCookies, Cookies } from 'react-cookie';
import { instanceOf } from 'prop-types';
import DanteLogo from '../DanteLogo.png';

const styles = ({
  auth: {
    position: 'absolute',
    margin: 'auto',
    top: 0,
    right: 0,
    bottom: 0,
    left: 0,
    width: '500px',
    height: '300px',
    backgroundColor: '#ccc',
    borderRadius: '3px',
  },
});


class Authentication extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '', password: '', redirect: false,
    };

    this.handleChangeEmail = this.handleChangeEmail.bind(this);
    this.handleChangePassword = this.handleChangePassword.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChangeEmail(event) {
    this.setState({ email: event.target.value });
  }

  handleChangePassword(event) {
    this.setState({ password: event.target.value });
  }

  handleSubmit(event) {
    const { email, password } = this.state;
    const { cookies } = this.props;
    event.preventDefault();
    if (email === '' || password === '') {
      this.setState({ errorMessage: 'Must supply credentials.' });
      return;
    }
    fetch('https://danteapi.chaz.pro/auth', {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email, password,
      }),
    }).then(response => response.json())
      .then((responseJson) => {
        if (responseJson.status === 'failed') {
          this.setState({ redirect: false, errorMessage: responseJson.message });
          return;
        }
        cookies.set('token', responseJson.token);
        this.setState({ redirect: true });
      })
      .catch(error => this.setState({ errorMessage: error.message }));
  }

  render() {
    const {
      email, password, errorMessage, redirect,
    } = this.state;
    if (redirect === true) {
      return (<Redirect to="/" />);
    }
    return (
      <div>
        <img
          src={DanteLogo}
          alt="DanteProject"
          style={{ maxWidth: '250px', height: '250px', margin: 'auto' }}
        />
        <div className="Auth" style={styles.auth}>
          <Typography variant="h4" gutterBottom>
           Dante Project
          </Typography>
          <Typography variant="subtitle1" gutterBottom>
            Auth
          </Typography>
          <form onSubmit={this.handleSubmit}>
            <TextField label="E-mail" type="text" value={email} onChange={this.handleChangeEmail} />
            <br />
            <TextField label="Password" type="password" value={password} onChange={this.handleChangePassword} />
            <br />
            <br />
            <Button variant="contained" color="secondary" type="submit">
              Submit
            </Button>
            <Button variant="contained" color="secondary" href="/register">Register</Button>
          </form>
          { errorMessage && <Typography variant="h5">{errorMessage}</Typography>}
        </div>
      </div>
    );
  }
}

Authentication.propTypes = {
  cookies: instanceOf(Cookies).isRequired,
};

export default withCookies(Authentication);
