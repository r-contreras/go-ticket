import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

const TextLink = styled(Link)`
    text-decoration: none;
    color: white;

    &:focus, &:hover, &:visited, &:link, &:active {
        text-decoration: none;
    } 
`

export default (props) => <TextLink {...props} />;