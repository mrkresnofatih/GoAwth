import { css } from '@emotion/css'
import React from 'react'

const AppButton = ({text, onClick, style}) => {
  return (
    <label className={appButtonStyle.body} onClick={onClick} style={style}>{text}</label>
  )
}

export default AppButton

const appButtonStyle = {
    body: css`
        padding: 12px 18px;
        background-color: #4CAF50;
        display: flex;
        justify-content: center;
        align-items: center;

        :hover {
            cursor: pointer;
            opacity: 0.9;
        }
    `
}