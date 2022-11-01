import { css } from '@emotion/css'
import { useRouter } from 'next/router'
import React from 'react'

const RouterLink = ({href, children}) => {
    const router = useRouter()

    const handleClick = (e) => {
        e.preventDefault()
        router.push(href)
    }

    return (
        <label className={routerLinkStyles.body} onClick={handleClick}>{children}</label>
    )
}

export default RouterLink

const routerLinkStyles = {
    body: css`
        display: flex;
    `
}