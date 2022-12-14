import { css } from '@emotion/css'
import React from 'react'

const TextInput = ({placeholder, password, style, onChange, value}) => {
    return (
        <input 
          placeholder={placeholder} 
          className={textInputStyle.body}
          type={password ? "password" : "text"}
          onChange={onChange}
          style={style}
          value={value}
        />
      )
}

export default TextInput

const textInputStyle = {
    body: css`
        padding: 12px;
        font-size: 16px;
        background-color: rgb(25,25,25);
        border: 1px solid white;
        border-radius: 6px;
        margin-top: 20px;

        ::-ms-reveal {
            filter: invert(100%);
        }
    `
}