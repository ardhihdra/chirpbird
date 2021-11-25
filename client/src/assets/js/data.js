import img2 from '../img/chart.png';

const fetchPosts = () => {
    return [
        {
            id: 1,
            title: "Ini Judul",
            highlight: img2,
            description: "Every human environment is embedded with its own intellectual",
            date: "Fri 12 March 2021",
            situation: "lagi cape",
            text: `<p>The export statement is used when creating JavaScript modules to export live bindings to functions, objects, or primitive values from the module so they can be used by other programs with the import statement. Bindings that are exported can still be modified locally; when imported, although they can only be read by the importing module the value updates whenever it is updated by the exporting module.</p>
            <p>The export statement is used when creating JavaScript modules to export live bindings to functions, objects, or primitive values from the module so they can be used by other programs with the import statement. Bindings that are exported can still be modified locally; when imported, although they can only be read by the importing module the value updates whenever it is updated by the exporting module.</p>
            <p>The export statement is used when creating JavaScript modules to export live bindings to functions, objects, or primitive values from the module so they can be used by other programs with the import statement. Bindings that are exported can still be modified locally; when imported, although they can only be read by the importing module the value updates whenever it is updated by the exporting module.</p>
            <p>The export statement is used when creating JavaScript modules to export live bindings to functions, objects, or primitive values from the module so they can be used by other programs with the import statement. Bindings that are exported can still be modified locally; when imported, although they can only be read by the importing module the value updates whenever it is updated by the exporting module.</p>
            `
        },
        {
            id: 2,
            title: "Ini Lainya",
            highlight: img2,
            description: "Every human is crazy with its own intellectual",
            date: "Fri 14 March 2021",
            situation: "segar"
        },
        {
            id: 3,
            title: "Ini Gatau",
            highlight: img2,
            description: "Every human is unique with its own intellectual",
            date: "Fri 06 April 2021",
            situation: "lagi ngelamun"
        }
    ];
}

export {
    fetchPosts
}