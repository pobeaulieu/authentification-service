import React, { useEffect, useState } from "react";
import PropTypes from 'prop-types';

interface Client {
    id: number;
    firstName: string;
    lastName: string;
    businessType: number;
}

const Business = (props: any) => {
    const [clients, setClients] = useState<Client[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        async function fetchClients() {
            const response = await fetch('http://localhost:8000/api/businessclients', {
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
            const data = await response.json();
            setClients(data);
            setLoading(false);
        }
        if(clients.length === 0) {
            fetchClients();
        }
    }, []);


    if (props.loggedUser.businessRole !== 1 && props.loggedUser.adminRole !== 1) {
        return (
            <div>
                Vous n'êtes pas autorisé à accéder à cette page
            </div>
        );
    }

    if (loading) {
        return <div><h2>Clients Affaires</h2>Loading...</div>;
    }

    return (
        <div>
            <h2>Clients Affaires</h2>
            <table className="type-table">
                <thead>
                <tr>
                    <th>Prénom</th>
                    <th>Nom</th>
                    <th>Type</th>
                </tr>
                </thead>
                <tbody>
                {clients.map((client) => (
                    <tr key={client.id}>
                    <td>{client.firstName}</td>
                    <td>{client.lastName}</td>
                    { client.businessType === 1 && <td>BUSINESS</td>}
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

Business.propTypes = {
    loggedUser: PropTypes.shape({
        id: PropTypes.number,
        name: PropTypes.string,
        email: PropTypes.string,
        adminRole: PropTypes.number,
        businessRole: PropTypes.number,
    }).isRequired,
};

export default Business;