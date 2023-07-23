import React from 'react';
import { Typography, Card, CardContent, Button } from '@mui/material';
import { format } from 'date-fns';
import Cookies from 'universal-cookie';

const Comment = ({ comment, onDeleteComment }) => {
  const { id, userid, content, createdat } = comment;

  const handleDelete = async () => {
    const cookies = new Cookies();
    const token = cookies.get('user');

    try {
      const response = await fetch(`http://localhost:9001/messages/${id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        onDeleteComment(id);
        window.location.reload(); // Actualizar la página
      } else {
        console.error('Error deleting comment:', response.status);
      }
    } catch (error) {
      console.error('Error deleting comment:', error);
    }
  };

  return (
    <Card>
      <CardContent>
        <Typography variant="subtitle1">Usuario: {userid}</Typography>
        <Typography variant="body1">{content}</Typography>
        <Typography variant="caption">Publicado el {createdat}</Typography>
        <Button variant="outlined" color="error" onClick={handleDelete}>
          Eliminar
        </Button>
      </CardContent>
    </Card>
  );
};

export default Comment;


export function CommentList({ comments , onDeleteComment}) {
    const handleDeleteComment = (commentId) => {
        onDeleteComment(commentId);
      };
    
    return (
      <div>
        {comments ? (
          comments.map((comment) => (
            <Comment key={comment.id} comment={comment} onDeleteComment={handleDeleteComment}/>
          ))
        ) : (
          <p>No hay comentarios</p>
        )}
      </div>
    );
  }