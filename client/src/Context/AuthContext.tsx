import {createContext, ReactNode, useContext, useState} from 'react';
import {Admin, AuthenticatedUser, Customer, User} from "../types.tsx";


 interface AuthContextProps {
    user: AuthenticatedUser | null;
    login: (userDetails: User, additionalInfo: Customer | Admin) => void;
    logout: () => void;
}
const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};
export const AuthProvider: React.FC<{children: ReactNode}> = ({ children }) => {
    const [user, setUser] = useState<AuthenticatedUser | null>(null);
    const login = (userDetails: User, additionalInfo: Customer | Admin) => {
        setUser({
            type: 'Customer' in additionalInfo ? 'Customer' : 'Admin',
            details: userDetails,
            additionalInfo: additionalInfo,
        });
    };

    const logout = () => setUser(null);

    const value = { user, login, logout };
    return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
