import React from 'react';
import { Redirect } from 'react-router';
import { instanceOf } from 'prop-types';
import { Cookies, withCookies } from 'react-cookie';

const Logout = (props) => {
  const { cookies } = props;

  cookies.remove('token');
  return (
    <Redirect to="/auth" />
  );
};

export default withCookies(Logout);

Logout.propTypes = {
  cookies: instanceOf(Cookies).isRequired,
};
