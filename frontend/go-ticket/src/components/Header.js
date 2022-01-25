import React from "react";

function Header(props) {
    return(
        <div>
          <h1 className="text-4xl font-bold text-center mr-20" > 
            {props.text}
          </h1>
        </div>
    )
}

export default Header;