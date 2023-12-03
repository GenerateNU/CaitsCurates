import Check from "../../images/check.svg";
import Error from "../../images/error.svg";

export type ValidationAlertProps = {
    fieldText: string;
    validationText: string;
    validationRule: RegExp;
    className?: string;
}

const ValidationAlert = ({ fieldText, validationText, validationRule, className }: ValidationAlertProps) => {
    return (
        <div className={`flex items-center ${className}`}>
            <img className="w-4 h-4 mr-1" src={fieldText.match(validationRule) ? Check : Error} />
            <span className="text-base"> {validationText} </span>
        </div>
    )
}

export default ValidationAlert;