import styles from '../Notes.module.css';
import PocketBase from 'pocketbase';

async function getNote(noteId: string) {
  const db = new PocketBase('http://127.0.0.1:8090');
  const authData = await db.admins.authWithPassword('joebthomas4@gmail.com', 'joethomas55');
  const res = await fetch(
    `http://127.0.0.1:8090/api/collections/posts/records/${noteId}`,
    {
      next: { revalidate: 10 },
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${authData.token}`
      }
    }
  );
  const data = await res.json();
  return data;
}

export default async function NotePage({ params }: any) {
  const note = await getNote(params.id);
  console.log("here!")

  return (
    <div>
      <h1>notes/{note.id}</h1>
      <div className={styles.note}>
        <h3>{note.title}</h3>
        <h5>{note.content}</h5>
        <p>{note.created}</p>
      </div>
    </div>
  );
}
