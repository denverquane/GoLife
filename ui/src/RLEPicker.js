import React, {useEffect, useRef} from "react";
import {makeStyles} from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Popover from '@material-ui/core/Popover';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme) => ({
    button: {
        display: 'block',
        marginTop: theme.spacing(2),
    },
    formControl: {
        margin: theme.spacing(1),
        minWidth: 120,
    },
    typography: {
        padding: theme.spacing(2),
    },
}));

export default function RLEPicker(props) {
    const classes = useStyles();
    const [rleName, setRLEName] = React.useState('');
    const [rle, setRLE] = React.useState(null);
    const [rleScale, setRLEScale] = React.useState(1);
    const [open, setOpen] = React.useState(false);
    const [anchorEl, setAnchorEl] = React.useState(null);

    const handleChange = (event) => {
        setRLEName(event.target.value);
        props.onSelect(event.target.value);
        props.rles.forEach(rle => {
            if (rle.getName() === event.target.value) {
                setRLE(rle);
            }
        })
    };

    const handleClose = () => {
        setOpen(false);
    };

    const handleOpen = () => {
        setOpen(true);
    };

    const handleClick = (event) => {
        setAnchorEl(event.currentTarget);
    };
    const handleClosePopover = () => {
        setAnchorEl(null);
    };

    const draw = (ctx, height, width, data, color, scale) => {
        let r = (color >> 24) & 0xFF;
        let g = (color >> 16) & 0xFF;
        let b = (color >> 8) & 0xFF;
        ctx.fillStyle = 'rgb(' + r + ',' + g + ',' + b + ')'
        //ctx.fillStyle = '#000000'
        let idx = 0;
        for (let y = 0; y < height; y++) {
            for (let x = 0; x < width; x++) {
                if (data[idx]) {
                    ctx.fillRect(x*scale, y*scale, scale, scale)
                }
                idx++;
            }
        }
    }

    const canvasRef = useRef(null)

    useEffect(() => {
        if (rle && canvasRef.current) {
            const canvas = canvasRef.current
            const context = canvas.getContext('2d')

            //Our draw come here
            draw(context, rle.getHeight(), rle.getWidth(), rle.getData(), props.color, rleScale)
        }
    }, [draw])

    const openPopover = Boolean(anchorEl);
    const id = open ? 'simple-popover' : undefined;

    return (
        <div style={{display: "flex", flexDirection: "column"}}>
            <FormControl className={classes.formControl}>
                <InputLabel id="demo-controlled-open-select-label">Pattern</InputLabel>
                <Select
                    labelId="demo-controlled-open-select-label"
                    id="demo-controlled-open-select"
                    open={open}
                    onClose={handleClose}
                    onOpen={handleOpen}
                    value={rleName}
                    onChange={handleChange}
                >
                    <MenuItem value="">
                        <em>None</em>
                    </MenuItem>
                    {props.rles.map((item, i) => {
                        return <MenuItem value={item.getName()}>{item.getName()}</MenuItem>
                    })}
                </Select>
            </FormControl>
            <Button aria-describedby={id} variant="contained" color="primary" onClick={handleClick}
                    disabled={rleName === ""}>
                View Pattern Info
            </Button>

            <Popover
                id={id}
                open={openPopover}
                anchorEl={anchorEl}
                onClose={handleClosePopover}
                anchorOrigin={{
                    vertical: 'bottom',
                    horizontal: 'center',
                }}
                transformOrigin={{
                    vertical: 'top',
                    horizontal: 'center',
                }}
            >
                {rle ?
                    <div>
                        <Typography
                            className={classes.typography}>Width: {rle.getWidth()} Height: {rle.getHeight()}</Typography>
                        {rle.getWidth() < 800 && rle.getHeight() < 450 ?
                            <canvas width={rle.getWidth()} height={rle.getHeight()} ref={canvasRef}></canvas> :
                            <Typography className={classes.typography}>Pattern is too large to preview!</Typography>}
                    </div>
                    : <div/>}

            </Popover>
        </div>
    )

}