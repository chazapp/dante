import React, { Component } from 'react';
import { withCookies, Cookies } from 'react-cookie';
import PropTypes from 'prop-types';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import store from '../redux/store';
import HeaderAppBar from '../common/HeaderAppBar';
import { fetchSources } from '../redux/actions';
import SourcesForm from '../common/SourcesForm';
import SourcesList from '../common/SourcesList';

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
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    email: state.profile.email,
    error: state.error,
  };
}

Home.propTypes = {
  cookies: PropTypes.instanceOf(Cookies).isRequired,
  sources: PropTypes.arrayOf(PropTypes.shape()).isRequired,
};

export default withCookies(connect(mapStateToProps)(Home));
