import React from "react";
import caitLogo from "../../images/cait_logo.svg";

const Footer = () => {
    return (
        <div className="flex h-72 py-14 px-20 bg-beige">
            <div className="flex flex-col gap-y-3">
                <img src={caitLogo} className="w-16 h-16" />
                {/* Non-functional */}
                <div className="font-proxima font-bold text-espresso text-sm"> Request </div>
                <div className="font-proxima font-bold text-espresso text-sm"> Shop </div>
                <div className="font-proxima font-bold text-espresso text-sm"> About </div>
                <div className="font-proxima font-bold text-rose text-sm"> Terms and Privacy Policy </div>

            </div>

        </div>
    )
};

export default Footer;