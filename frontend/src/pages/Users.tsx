import React, { useEffect, useState } from "react";
import PropTypes from 'prop-types';
import moment from 'moment';
import Register from "../components/Register";

interface User {
    id: number;
    name: string;
    email: string;
    adminRole: number;
    businessRole: number;
    residentialRole: number;
    lastModified: string;
    lastLoginAttempt: string;
    loginAttemptCount: number;
    blocked: number;
}

const Users = (props: any) => {

    const [users, setUsers] = useState<User[]>([]);
    const [loading, setLoading] = useState(true);
    const [visibleUserActions, setVisibleUserActions] = useState<number |null>(null);
    const [isModalOpen, setIsModalOpen] = useState(false);

    useEffect(() => {
        async function fetchClients() {
            const response = await fetch('http://localhost:8000/api/users', {
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
            const data = await response.json();
            setUsers(data);
            setLoading(false);
        }
        if(users.length === 0) {
            fetchClients();
        }
    }, []);

    const handleToggleUserDiv = (userId: number) => {
        setVisibleUserActions((prevVisibleUserDiv: number | null) => {
            if (prevVisibleUserDiv === userId) {
                return null;
            } else {
                return userId;
            }
        });
    };

    const handleRemoveRole = async (userId: number, role: string) => {
        console.log("Remove role " + role + " from user " + userId);
    }

    const handleAddRole = async (userId: number, role: string) => {
        console.log("Add role " + role + " to user " + userId);
    }

    const handleUnblockUser = async (userId: number) => {
        const response = await fetch('http://localhost:8000/api/unblock', {
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({userId: userId}),
        });
        const data = await response.json();
        setUsers(data);
        setLoading(false);
    }

    const handleResetPassword = async (userId: number) => {
        console.log("Reset password for user " + userId);
    }

    function convertDate(date: string) {
        const dt = moment("2023-03-25T11:46:47.496262-04:00");
        return dt.format("Do MMMM YYYY, h:mm:ss A z");
    }

    // const handleOpenModal = () => {
    //     setIsModalOpen(true);
    // };
    
    // const handleCloseModal = () => {
    //     setIsModalOpen(false);
    // };
      
    // function Modal(props: any) {
    //     return (
    //         <div className="modal">
    //         <div className="modal-content">
    //             <h2>Créer un utilisateur</h2>
    //             <Register onSubmit={props.onClose}/>
    //             <button onClick={props.onClose}>Fermer</button>
    //         </div>
    //         </div>
    //     );
    // }

    if (props.loggedUser.residentialRole !== 1) {
        return (
            <div>
                Vous n'êtes pas autorisé à accéder à cette page
            </div>
        );
    }

    if (loading) {
        return <div>
            <h2>Utilisateurs</h2>
            Loading...</div>;
    }

    return (
        <div>
            <h2>Utilisateurs</h2>
            {/* <button onClick={handleOpenModal}>Créer utilisateur</button>
            {isModalOpen && <Modal onClose={handleCloseModal}/>} */}
            <table className="type-table">
                <thead>
                <tr>
                    <th>Id</th>
                    <th>Nom d'utilisateur</th>
                    <th>Rôles</th>
                    <th>Bloqué</th>
                    <th>Mot de passe modifié</th>
                    <th>Dernière tentative de connexion</th>
                    <th>Total tentatives échouées</th>
                    <th>Actions</th>
                </tr>
                </thead>
                <tbody>
                {users.map((user) => (
                    <tr key={user.id}>
                        <td>{user.id}</td>
                        <td>{user.name}</td>
                        <td>
                            {user.adminRole === 1 && 
                                <div>
                                    <span>Admin</span>
                                    {/* <button onClick={() => handleRemoveRole(user.id, "admin")} style={{marginLeft: "10px", fontSize: "0.8em"}}>
                                        Remove
                                    </button> */}
                                </div>
                            }
                            {user.businessRole === 1 && 
                                <div>
                                    <span>Business</span>
                                    {/* <button onClick={() => handleRemoveRole(user.id, "business")} style={{marginLeft: "10px", fontSize: "0.8em"}}>
                                        Remove
                                    </button> */}
                                </div>
                            }
                            {user.residentialRole === 1 &&
                                <div>
                                    <span>Residential</span>
                                    {/* <button onClick={() => handleRemoveRole(user.id, "residential")} style={{marginLeft: "10px", fontSize: "0.8em"}}>
                                        Remove
                                    </button> */}
                                </div>
                            }
                        </td>
                        <td>
                            {user.blocked === 1 && <span><b>X</b><button onClick={() => handleUnblockUser(user.id)} style={{marginLeft: "10px", fontSize: "0.8em"}}>Débloquer</button></span>}
                        </td>
                        <td>{convertDate(user.lastModified)}</td>
                        <td>
                            {user.lastLoginAttempt !== "0000-12-31T19:00:00-05:00" && <span>{convertDate(user.lastLoginAttempt)}</span>}
                        </td>
                        <td>{user.loginAttemptCount}</td>
                        <td>
                            <button onClick={() => handleToggleUserDiv(user.id)}>
                                Actions
                            </button>
                            {visibleUserActions === user.id && (
                                <div className="hidden-div">
                                    {/* <button onClick={() => handleAddRole(user.id, "admin")}>Ajouter rôle admin</button><br />
                                    <button onClick={() => handleAddRole(user.id, "business")}>Ajouter rôle business</button><br />
                                    <button onClick={() => handleAddRole(user.id, "residential")}>Ajouter rôle residential</button><br /> */}
                                    <button onClick={() => handleResetPassword(user.id)}>Réinitialiser mot de passe</button>
                                </div>
                            )}
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

Users.propTypes = {
    loggedUser: PropTypes.shape({
        id: PropTypes.number,
        name: PropTypes.string,
        email: PropTypes.string,
        adminRole: PropTypes.number,
    }).isRequired,
};

export default Users;