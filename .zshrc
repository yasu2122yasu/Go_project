# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('/Users/yasushitamura/opt/anaconda3/bin/conda' 'shell.zsh' 'hook' 2> /dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "/Users/yasushitamura/opt/anaconda3/etc/profile.d/conda.sh" ]; then
        . "/Users/yasushitamura/opt/anaconda3/etc/profile.d/conda.sh"
    else
        export PATH="/Users/yasushitamura/opt/anaconda3/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda initialize <<<

export PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
export PATH=$PATH:/usr/local/go/bin 


# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('/Users/yasushitamura/opt/anaconda3/bin/conda' 'shell.zsh' 'hook' 2> /dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "/Users/yasushitamura/opt/anaconda3/etc/profile.d/conda.sh" ]; then
        . "/Users/yasushitamura/opt/anaconda3/etc/profile.d/conda.sh"
    else
        export PATH="/Users/yasushitamura/opt/anaconda3/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda initialize <<<

export PATH=/usr/local/bin:$PATH

#Additional PATH
export PATH=$PATH:/sbin
eval "$(direnv hook zsh)"
export PATH=$HOME/go/bin:$PATH
eval "$(direnv hook zsh)"