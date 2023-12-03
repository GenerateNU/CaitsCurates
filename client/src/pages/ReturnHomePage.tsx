import HomePrompt from "../components/Home/HomePrompt";
import CaitThumbnail1 from "../images/cait_thumbnail_1.svg";
import CaitThumbnail2 from "../images/cait_thumbnail_2.svg";
import CaitThumbnail3 from "../images/cait_thumbnail_3.svg";
import Or from "../images/or.svg";
import Navbar from "../components/Home/NavBarUpdated";

const ReturnHomePage = () => {

    const handleShoppingClick = () => {
        // handle click
    }

    const handleBrowsingClick = () => {
        // handle click
    }

    const handleRequestsClick = () => {
        // handle click
    }
    
    return (
        <div>
            <Navbar />
            <div className="flex flex-col bg-eggshell">
                <div className="flex h-[38rem] items-center overflow-hidden">
                    <HomePrompt
                        className="w-[512px] pl-40 pr-28 shrink-0 lg:w-1/2"
                        titleText="Gift-Giving With Love"
                        subText="Give your loved ones the perfect gift with Cait's Curates personalized gifiting."
                        buttonText="Start Shopping!"
                        onClick={handleShoppingClick}
                    />
                    <img className="pl-20" src={CaitThumbnail1} />
                </div>
                <div className="grid grid-cols-1 bg-beige px-[5rem] py-[5rem] md:px-[13rem] lg:grid-cols-3">
                    <HomePrompt
                        className="justify-center"
                        textStyle="text-[1.5rem]"
                        imageStyle="shrink-0 self-center lg:self-start"
                        image={CaitThumbnail2}
                        titleText="Browse Cait's Picks"
                        subText="Browse Cait's curated collections of gifts for all occasions and recipients!"
                        buttonText="Start Browsing"
                        onClick={handleBrowsingClick}
                    />
                    <div className="row-start-1 flex flex-col items-center px-12 lg:row-start-auto">
                        <div className="text-4xl font-seasons font-bold">
                            {"How it Works"}
                        </div>
                        <img className="h-0 w-0 mt-20 sm:w-32 lg:h-32" src={Or} />
                        <div className="relative bottom-[4.8rem] text-[0] text-white font-seasons font-bold lg:text-3xl"> {"or"} </div>
                    </div>
                    <HomePrompt
                        className="justify-center mt-20 lg:mt-0"
                        textStyle="text-[1.5rem]"
                        imageStyle="self-center lg:self-start"
                        image={CaitThumbnail3}
                        titleText="Let Cait Do It For You!"
                        subText="Fill out a simple form and let Cait curate the perfect gift for your occasion and recipient!"
                        buttonText="Submit a Request"
                        onClick={handleRequestsClick}
                    />
                </div>

            </div>
        </div>
    )
}

export default ReturnHomePage;