import { css } from '@emotion/css'
import React from 'react'
import PageHeader from './PageHeader'
import SeoHead from './SeoHead'

const PageTemplate = ({seoHead, children}) => {
  return (
    <div className={pageTemplateStyles.body}>
        <SeoHead title={seoHead.title} description={seoHead.description} />
        <PageHeader title={"Home"} />
        {children}
    </div>
  )
}

export default PageTemplate

const pageTemplateStyles = {
    body: css`
        background-color: black;
        display: flex;
        flex-direction: column;
        height: 100vh;
        width: 100vw;
    `
}