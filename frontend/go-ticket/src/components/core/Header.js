import React from "react";

function Header({ text }) {
  return (
    <div>
      <h1 className="display-3" >
        {text}
      </h1>
    </div>
  )
}

export default Header;