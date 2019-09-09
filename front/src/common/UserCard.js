import React from 'react';
import PropTypes from 'prop-types';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import DefaultPicture from '../DefaultUser.png';

const UserCard = (props) => {
  const { handle, email, picture } = props;

  const userImage = () => {
    if (picture) {
      return (
        <img
          src={`data:image/jpeg;base64,${picture}`}
          alt="ProfilePicture"
          style={{ width: '200px', height: '200px' }}
        />
      );
    }
    return (
      <img
        src={DefaultPicture}
        style={{ width: '200px', height: '200px' }}
        alt="UserPicture"
      />
    );
  };

  return (
    <div style={{ float: 'left' }}>
      {userImage()}
      <Typography variant="h5">
        {'@'}
        {handle}
        <br />
        {email}
      </Typography>
      <Button variant="contained" href="/edit">
        Edit
      </Button>
    </div>
  );
};

UserCard.propTypes = {
  handle: PropTypes.string.isRequired,
  email: PropTypes.string.isRequired,
  picture: PropTypes.string.isRequired,
};

export default UserCard;
