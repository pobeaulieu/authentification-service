import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import { Tab, Tabs } from 'react-bootstrap';

const Administration = (props: any) => {
    // useEffect(() => {
    //     async function fetchClients() {
    //         const response = await fetch('http://localhost:8000/api/users', {
    //             headers: {'Content-Type': 'application/json'},
    //             credentials: 'include',
    //         });
    //         const data = await response.json();
    //     }
    // }, []);

    const handleConfigConnection = (event: any) => {
        event.preventDefault();
    }

    if (props.loggedUser.adminRole !== 1) {
        return (
            <div>
                Vous n'êtes pas autorisé à accéder à cette page
            </div>
        );
    }

    return (
        <div>
            <h2>{'Administration'}</h2>
            <Tabs defaultActiveKey="connection" className="mb-3" >
                <Tab eventKey="connection" title="Configuration de connexion">
                    <form onSubmit={handleConfigConnection}>
                        <div className="form-group">
                            <label htmlFor="inputNbConnexion">Nombre de connexion</label>
                            <input type="number" className="form-control" id="inputNbConnexion" value={5} />
                        </div>
                    </form>
                </Tab>
                <Tab eventKey="complexPxd" title="Configuration de complexité du mot de passe">
                    <p>tab2</p>
                </Tab>
                <Tab eventKey="modifyPwd" title="Configuration de Modification du mot de passe">
                    <p>tab3</p> 
                </Tab>
            </Tabs>
        </div>
      );
};

Administration.propTypes = {
    loggedUser: PropTypes.shape({
        id: PropTypes.number,
        name: PropTypes.string,
        email: PropTypes.string,
        adminRole: PropTypes.number,
    }).isRequired,
};

export default Administration;