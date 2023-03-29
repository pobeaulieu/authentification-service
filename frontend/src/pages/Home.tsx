import React from "react";
import PropTypes from 'prop-types';

const Home = (props: any) => {
    return (
        <div>
            {props.loggedUser.name ? 'Bonjour ' + props.loggedUser.name : 'Vous n\'êtes pas connecté'}
        </div>
    );
};

Home.propTypes = {
    loggedUser: PropTypes.shape({
        id: PropTypes.number,
        name: PropTypes.string,
        email: PropTypes.string,
    }).isRequired,
};

export default Home;