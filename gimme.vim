function! s:Gimme(type)
  let saved = @@
  if a:type ==# 'V'
    normal! `<V`>y
  elseif a:type ==# 'v'
    normal! `<v`>y
  elseif a:type ==# 'char'
    normal! `[v`]y
  else
    return
  endif
  " remove trailing newline
  "
  "   https://stackoverflow.com/a/6228454
  let @@ = substitute(strtrans(@@), '\^@', '', 'gimme')
  call system('gimme ' . shellescape(@@))
  let @@ = saved
endfunction

vnoremap <F2> :<c-u>call <SID>Gimme(visualmode())<cr>
nnoremap <F2> :set operatorfunc=<SID>Gimme<cr>g@
nnoremap <F2><F2> :silent !g <c-r><c-w><cr>
