// Your CSS as text
var styles = `
    .test-inline-style div { 
        font-family: Georgia,Cambria,"Times New Roman",Times,serif;
        margin: 26px auto 0 auto;
        max-width: 650px;
    }
`;

var styleSheet = document.createElement("style");
styleSheet.type = "text/css";
styleSheet.innerText = styles;
document.head.appendChild(styleSheet);
