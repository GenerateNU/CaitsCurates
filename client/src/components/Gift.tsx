import React from "react";
import { Gift } from "../types";

const Gift: React.FC<Gift> = ({ Name, Link }: Gift) => {
    return (
        <div>
            <p className="">{Name}</p>
            <a className="text-blue-800" href={Link} target="_blank" rel="noreferrer">
                View product
            </a>
        </div>
    );
};

export default Gift;