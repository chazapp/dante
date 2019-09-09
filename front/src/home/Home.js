import React, { Component } from 'react';
import { withCookies, Cookies } from 'react-cookie';
import PropTypes from 'prop-types';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import HeaderAppBar from '../common/HeaderAppBar';

/* const styles = {
  UserCard: {
    float: 'left',
  },
  Content: {
  },
}; */

class Home extends Component {
  constructor(props) {
    super(props);
    const { cookies } = props;
    let token = '';
    if (cookies.get('token') !== null) {
      token = cookies.get('token');
    }
    this.state = {
      token,
    };
  }

  componentDidMount() {
    const { token } = this.state;
  }

  render() {
    const {
      token,
    } = this.state;
    const { sources } = this.props;
    if (token === undefined) {
      return (<Redirect to="/auth" />);
    }
    return (
      <div>
        <HeaderAppBar />
        <p>Coming soon<br/>Dante Front-End</p>
        <p>{token}</p>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
  };
}

Home.propTypes = {
  cookies: PropTypes.instanceOf(Cookies).isRequired,
};

export default withCookies(connect(mapStateToProps)(Home));
