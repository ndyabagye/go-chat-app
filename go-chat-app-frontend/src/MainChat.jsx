import {useNavigate, useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import ChannelsList from "./ChannelsList"
import MessagesPanel from "./MessagesPanel"

const MainChat = ( ) =>{
    const {channelId} = useParams();
    const navigate = useNavigate();
    const [selectedChannel, setSelectedChannel] = useState(parseInt(channelId) || null);

    // if a component loads with channel ID in the URL, set it as the selected channel
    useEffect(()=>{
        if (selectedChannel){
            navigate(`/chat/${selectedChannel.id}`);
        }
    },[selectedChannel, navigate]);

    const handleChannelSelect = (channelId) =>{
        setSelectedChannel(channelId);
    }

    return (
        <div className="flex h-screen">
            <div className="w-1/4 border-r">
                <ChannelsList
                    selectedChannel={selectedChannel}
                    setSelectedChannel={handleChannelSelect}
                />
            </div>
            <div className="w-3/4">
                <MessagesPanel selectedChannel={selectedChannel}/>
            </div>
        </div>
    )
}

export default MainChat;